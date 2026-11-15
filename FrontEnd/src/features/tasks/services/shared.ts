import type { TaskItem } from "../types/type";

export const delay = (ms = 400) =>
  new Promise<void>((resolve) => setTimeout(resolve, ms));

const DAY = 86_400_000;

export const isoFromNow = (offsetDays: number) =>
  new Date(Date.now() + offsetDays * DAY).toISOString();

/**
 * Satu-satunya tempat aturan transisi status ditulis.
 * Dipakai bersama oleh service server (baca) dan aksi klien (ubah).
 */

export const assertCanStart = (task: TaskItem) => {
  if (task.status !== "not_started") {
    throw new Error("Tugas ini sudah dimulai.");
  }
};

export const assertCanSubmit = (task: TaskItem) => {
  if (task.status === "submitted") {
    throw new Error("Tugas ini sedang menunggu review mentor.");
  }

  if (task.status === "completed") {
    throw new Error("Tugas ini sudah selesai dan disetujui mentor.");
  }
};
