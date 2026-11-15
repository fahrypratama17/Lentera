import CohortDetailLayoutContainer from "@/features/member-cohort/[slug]/container/CohortDetailLayoutContainer";

export default function CohortDetailLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: Promise<{ slug: string }>;
}) {
  return (
    <CohortDetailLayoutContainer params={params}>
      {children}
    </CohortDetailLayoutContainer>
  );
}