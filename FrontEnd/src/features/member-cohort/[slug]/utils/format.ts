const dateFormatter = new Intl.DateTimeFormat("id-ID", {
  day: "numeric",
  month: "long",
  year: "numeric",
  timeZone: "Asia/Jakarta",
});

export const formatDate = (iso: string) => dateFormatter.format(new Date(iso));

export const formatPeriod = (start: string, end: string) =>
  `${formatDate(start)} – ${formatDate(end)}`;