import { notFound, redirect } from "next/navigation";

import { getCohortDetail } from "../services/cohortDetailService";
import { buildCohortSlug, parseCohortSlug } from "../../utils/slugify";

interface CohortDetailRedirectContainerProps {
  params: Promise<{ slug: string }>;
}

const CohortDetailRedirectContainer = async ({
  params,
}: CohortDetailRedirectContainerProps) => {
  const { slug } = await params;

  const key = parseCohortSlug(slug);
  if (!key) notFound();

  const cohort = await getCohortDetail(key);
  if (!cohort) notFound();

  redirect(`/cohort/${buildCohortSlug(cohort.title, cohort.id)}/materi`);
};

export default CohortDetailRedirectContainer;
