export const slugify = (text: string) =>
  text
    .toLowerCase()
    .trim()
    .replace(/[^a-z0-9\s-]/g, "")
    .replace(/\s+/g, "-")
    .replace(/-+/g, "-");

export const buildCohortSlug = (title: string, id: string | number) =>
  `${slugify(title)}-${id}`;


export const parseCohortSlug = (slug: string) => {
  const decoded = decodeURIComponent(slug ?? "").trim();
  return decoded.length > 0 ? decoded : null;
};
