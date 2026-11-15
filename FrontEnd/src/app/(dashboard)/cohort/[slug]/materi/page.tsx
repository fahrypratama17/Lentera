import { Suspense } from "react";

import CohortMaterialContainer from "@/features/member-cohort/[slug]/container/CohortMaterialContainer";
import MaterialSkeleton from "@/features/member-cohort/[slug]/components/MaterialSkeleton";

async function MaterialSection({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const { slug } = await params;

  return <CohortMaterialContainer cohortId={slug} />;
}

export default function CohortMaterialPage({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  return (
    <Suspense fallback={<MaterialSkeleton />}>
      <MaterialSection params={params} />
    </Suspense>
  );
}