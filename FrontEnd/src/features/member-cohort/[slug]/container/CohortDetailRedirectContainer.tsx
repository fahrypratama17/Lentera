import { notFound, redirect } from "next/navigation";

import { getCohortDetail } from "../services/cohortDetailService";
import { buildCohortSlug, parseCohortSlug } from "../../utils/slugify";

const CohortDetailRedirectContainer = async ({ slug }: { slug: string }) => {
  const key = parseCohortSlug(slug);
  if (!key) notFound();

  const cohort = await getCohortDetail(key);
  if (!cohort) notFound();

  redirect(`/cohort/${buildCohortSlug(cohort.title, cohort.id)}/materi`);
};

export default CohortDetailRedirectContainer;
