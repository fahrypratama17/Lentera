import { notFound } from "next/navigation";

import MaterialList from "../components/MaterialList";
import { getCohortMaterials } from "../services/cohortDetailService";
import { parseCohortSlug } from "../../utils/slugify";

const CohortMaterialContainer = async ({ slug }: { slug: string }) => {
  const key = parseCohortSlug(slug);
  if (!key) notFound();
  const materials = await getCohortMaterials(key);

  return <MaterialList materials={materials} />;
};

export default CohortMaterialContainer;