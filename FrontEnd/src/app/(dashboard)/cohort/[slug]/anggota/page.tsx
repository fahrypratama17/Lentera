import { Suspense } from "react";

import CohortMemberContainer from "@/features/member-cohort/[slug]/container/CohortMemberContainer";
import MemberSkeleton from "@/features/member-cohort/[slug]/components/MemberSkeleton";

export default function CohortMemberPage({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  return (
    <Suspense fallback={<MemberSkeleton />}>
      <CohortMemberContainer params={params} />
    </Suspense>
  );
}