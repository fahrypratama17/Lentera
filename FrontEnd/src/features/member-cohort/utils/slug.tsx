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
  const id = slug.slice(slug.lastIndexOf("-") + 1);
  return /^\d+$/.test(id) ? id : null;
};