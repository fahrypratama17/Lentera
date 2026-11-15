export const slugify = (text: string) =>
  text
    .toLowerCase()
    .trim()
    .replace(/[^a-z0-9\s-]/g, "")
    .replace(/\s+/g, "-")
    .replace(/-+/g, "-");

export const buildCohortSlug = (title: string, id: string | number) => {
  const slug = slugify(title);
  const key = String(id);

  /**
   * Sebagian cohort memakai id yang sudah berupa slug judulnya
   * (mis. id `frontend-development`). Jangan digabung lagi agar URL
   * tidak menjadi `frontend-development-frontend-development`.
   */
  if (key === slug || key.startsWith(`${slug}-`)) return key;

  return `${slug}-${key}`;
};

export const parseCohortSlug = (slug: string) => {
  const decoded = decodeURIComponent(slug ?? "").trim();
  return decoded.length > 0 ? decoded : null;
};
