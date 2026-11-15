import CohortMaterialContainer from "@/features/member-cohort/[slug]/container/CohortMaterialContainer";

export default async function CohortMaterialPage({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const { slug } = await params;

  return <CohortMaterialContainer slug={slug} />;
}