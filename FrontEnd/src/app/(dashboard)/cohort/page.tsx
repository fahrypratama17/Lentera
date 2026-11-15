import CohortDetailRedirectContainer from "@/features/member-cohort/[slug]/container/CohortDetailRedirectContainer";

export default async function CohortDetailPage({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const { slug } = await params;
  return <CohortDetailRedirectContainer slug={slug} />;
}
