import {
  BookOpen,
  Target,
  UserCheck,
  BarChart3,
  UsersRound,
  ClipboardCheck,
  ChartNoAxesCombined,
  BookOpenCheck,
  MessagesSquare,
} from "lucide-react";
import type { LucideIcon } from "lucide-react";

type StepVariant = "tint" | "plain" | "active";

interface Step {
  n: number;
  icon: LucideIcon;
  title: string;
  text: string;
  variant: StepVariant;
}

export const NAV_LINKS = [
  { label: "Cara menggunakan", href: "#cara-menggunakan" },
  { label: "Manfaat", href: "#manfaat" },
  { label: "Tentang", href: "#tentang" },
];

export const STEPS: Step[] = [
  {
    n: 1,
    icon: UsersRound,
    title: "Buat atau Gabung Kelas",
    text: "Mulai dengan membuat kelas sebagai mentor atau bergabung sebagai anggota. Semua aktivitas belajar tersusun dalam satu ruang.",
    variant: "tint",
  },
  {
    n: 2,
    icon: BookOpen,
    title: "Bagikan Materi & Tugas",
    text: "Mentor membagikan materi dan tugas dengan mudah. Anggota bisa langsung tahu apa yang perlu dipelajari.",
    variant: "plain",
  },
  {
    n: 3,
    icon: ClipboardCheck,
    title: "Kerjakan & Kumpulkan",
    text: "Selesaikan tugas dan kumpulkan hasilnya langsung di kelas. Tidak perlu lagi mencari instruksi di grup chat.",
    variant: "plain",
  },
  {
    n: 4,
    icon: ChartNoAxesCombined,
    title: "Pantau Progres & Review",
    text: "Lihat perkembangan belajar, periksa tugas yang menunggu review, dan lanjutkan langkah berikutnya.",
    variant: "tint",
  },
] as const;

export const BENEFITS = [
  {
    icon: BookOpenCheck,
    title: "Semua Lebih Terorganisir",
    text: "Materi, tugas, dan pengumpulan tersimpan rapi dalam satu kelas. Tidak perlu lagi mencari informasi di antara tumpukan chat.",
  },
  {
    icon: ChartNoAxesCombined,
    title: "Progres Terlihat Jelas",
    text: "Pantau siapa yang sudah menyelesaikan tugas, siapa yang belum mulai, dan apa yang masih perlu ditinjau.",
  },
  {
    icon: MessagesSquare,
    title: "Review Jadi Lebih Mudah",
    text: "Mentor dapat meninjau hasil pekerjaan dengan lebih terarah, sementara anggota tahu apa yang perlu diperbaiki.",
  },
] as const;

export const STATS = [
  { label: "Role", num: "2", desc: "Mentor & Member, per cohort." },
  { label: "Status", num: "4", desc: "Belum mulai sampai Selesai." },
  { label: "Matriks", num: "1", desc: "Seluruh progres dalam satu tampilan." },
];

export const VARIANTS = {
  tint: {
    card: "border-secondary/10 bg-secondary-light",
    icon: "bg-white text-secondary",
    number: "text-secondary/50",
    title: "text-primary-text",
    text: "text-secondary-text",
    accent: "bg-secondary/50",
  },
  plain: {
    card: "border-border bg-background",
    icon: "bg-secondary-light text-secondary",
    number: "text-secondary/40",
    title: "text-primary-text",
    text: "text-secondary-text",
    accent: "bg-secondary/30",
  },
  active: {
    card: "border-primary bg-primary shadow-[0_18px_40px_-18px_rgba(31,80,109,0.55)]",
    icon: "bg-white/15 text-white",
    number: "text-white/40",
    title: "text-white",
    text: "text-white/75",
    accent: "bg-white/70",
  },
};
