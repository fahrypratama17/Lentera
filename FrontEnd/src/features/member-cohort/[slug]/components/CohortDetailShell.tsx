import { notFound } from "next/navigation";

import CohortDetailHeader from "./CohortDetailHeader";
import CohortTabs from "./CohortTabs";
import {
  getCohortDetail,
  getCohortMaterials,
} from "../services/cohortDetailService";
import { parseCohortSlug } from "../../utils/slugify";

const CohortDetailShell = async ({
  params,
}: {
  params: Promise<{ slug: string }>;
}) => {
  const { slug } = await params; 

  const key = parseCohortSlug(slug);
  if (!key) notFound();

  const [cohort, materials] = await Promise.all([
    getCohortDetail(key),
    getCohortMaterials(key),
    ]);
  if (!cohort) notFound();

  return (
    <>
      <CohortDetailHeader cohort={cohort} materials={materials} />
      <CohortTabs slug={slug} role={cohort.role} />
    </>
  );
};

export default CohortDetailShell;