import { cache } from "react";

import { mockCohorts } from "@/features/dashboard/data/data";
import type {
  CohortDetail,
  CohortMember,
  MaterialWithProgress,
} from "../types/type.ts"
import { buildCohortSlug } from "../../utils/slugify";

const wait = (ms = 500) => new Promise((resolve) => setTimeout(resolve, ms));

const DAY = 24 * 60 * 60 * 1000;
const iso = (offsetDays: number) =>
  new Date(Date.UTC(2026, 10, 1) + offsetDays * DAY).toISOString();


const findCohort = (slugOrId: string) =>
mockCohorts.find(
(cohort) =>
    String(cohort.id) === slugOrId ||
    buildCohortSlug(cohort.title, cohort.id) === slugOrId,
);

const materialTitles = [
  "Pengenalan & Persiapan",
  "Konsep Dasar",
  "Praktik Terarah",
  "Studi Kasus",
  "Proyek Mini",
  "Evaluasi Tengah",
  "Proyek Akhir",
  "Presentasi",
  "Refleksi",
];

const progressPattern = [
  { status: "completed", decision: "approved", note: "Kerja bagus, pemahamanmu sudah tepat." },
  { status: "completed", decision: "approved", note: "Sudah sesuai. Lanjutkan!" },
  { status: "in_progress", decision: "revision", note: "Tambahkan penjelasan pada bagian kesimpulan, lalu kumpulkan ulang." },
  { status: "submitted", decision: null, note: null },
] as const;

const memberNames = [
  "Aulia Rahma", "Bagas Wicaksono", "Callista Dewi", "Daffa Maulana",
  "Eka Prasetya", "Farah Nabila", "Gilang Ramadhan", "Hana Salsabila",
  "Ivan Kurniawan", "Jasmine Putri",
];

const toEmail = (name: string) =>
  `${name.toLowerCase().replace(/\s+/g, ".")}@mail.com`;


export const getCohortDetail = cache(
  async (slugOrId: string): Promise<CohortDetail | null> => {
    await wait();
    const cohort = findCohort(slugOrId);
    if (!cohort) return null;
    return { ...cohort,
      description: `Kelas ${cohort.title} membahas konsep dan praktik dari dasar hingga siap dipakai di proyek nyata, dibimbing langsung oleh ${cohort.mentorName}.`,
      startDate: iso(0),
      endDate: iso(91), };
  },
);


export const getCohortMaterials = cache(
  async (slugOrId: string): Promise<MaterialWithProgress[]> => {
    await wait();

    const cohort = findCohort(slugOrId);
    if (!cohort) return [];

    return Array.from({ length: cohort.materialCount }, (_, index) => {
      const pattern = progressPattern[index];

      return {
        id: `${slugOrId}-${index + 1}`,
        position: index + 1,
        title: materialTitles[index % materialTitles.length],
        deadline: iso((index + 1) * 7),
        status: pattern?.status ?? "not_started",
        reviewDecision: pattern?.decision ?? null,
        reviewNote: pattern?.note ?? null,
      };
    });
  },
);

export const getCohortMembers = cache(
  async (slugOrId: string): Promise<CohortMember[]> => {
    await wait(300);

    const cohort = findCohort(slugOrId);
    if (!cohort) return [];

    const mentor: CohortMember = {
      id: "mentor",
      name: cohort.mentorName,
      email: toEmail(cohort.mentorName),
      role: "mentor",
      joinedAt: iso(0),
      isMe: false,
    };

    const members: CohortMember[] = Array.from(
      { length: cohort.memberCount },
      (_, index) => {
        const round = Math.floor(index / memberNames.length);
        const name = `${memberNames[index % memberNames.length]}${
          round > 0 ? ` ${round + 1}` : ""
        }`;

        return {
          id: `member-${index + 1}`,
          name,
          email: toEmail(name),
          role: "member",
          joinedAt: iso(index % 10),
          isMe: index === 0,
        };
      },
    );

    return [mentor, ...members];
  },
);