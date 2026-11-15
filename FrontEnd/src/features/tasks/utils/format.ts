const dateFormatter = new Intl.DateTimeFormat("id-ID", {
  day: "numeric",
  month: "long",
  year: "numeric",
  timeZone: "Asia/Jakarta",
});

const dateTimeFormatter = new Intl.DateTimeFormat("id-ID", {
  day: "numeric",
  month: "short",
  year: "numeric",
  hour: "2-digit",
  minute: "2-digit",
  timeZone: "Asia/Jakarta",
});

export const formatDate = (iso: string) => dateFormatter.format(new Date(iso));

export const formatDateTime = (iso: string) =>
  dateTimeFormatter.format(new Date(iso));

const DAY = 86_400_000;

export const formatDueLabel = (
  deadline: string | null,
  now: number = Date.now(),
) => {
  if (!deadline) return "Tanpa tenggat";

  const days = Math.ceil((new Date(deadline).getTime() - now) / DAY);

  if (days < -1) return `Terlambat ${Math.abs(days)} hari`;
  if (days === -1) return "Terlambat 1 hari";
  if (days === 0) return "Tenggat hari ini";
  if (days === 1) return "Tenggat besok";
  if (days <= 7) return `${days} hari lagi`;

  return `Tenggat ${formatDate(deadline)}`;
};
