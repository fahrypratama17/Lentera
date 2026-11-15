import CohortMemberContainer from "@/features/member-cohort/[slug]/container/CohortMemberContainer";

export default async function CohortMemberPage({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const { slug } = await params;

  return <CohortMemberContainer slug={slug} />;
}