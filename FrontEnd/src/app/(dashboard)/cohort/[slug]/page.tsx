import { Suspense } from "react";

import CohortDetailRedirectContainer from "@/features/member-cohort/[slug]/container/CohortDetailRedirectContainer";
import CohortDetailSkeleton from "@/features/member-cohort/[slug]/components/CohortDetailSkeleton";

export default function CohortDetailPage({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  return (
    <Suspense fallback={<CohortDetailSkeleton />}>
      <CohortDetailRedirectContainer params={params} />
    </Suspense>
  );
}