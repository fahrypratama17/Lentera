import { Suspense } from "react";
import Link from "next/link";
import { ArrowLeft } from "lucide-react";

import CohortDetailShell from "../components/CohortDetailShell";
import CohortDetailSkeleton from "../components/CohortDetailSkeleton";

interface CohortDetailLayoutContainerProps {
  params: Promise<{ slug: string }>;
  children: React.ReactNode;
}

const CohortDetailLayoutContainer = ({
  params,
  children,
}: CohortDetailLayoutContainerProps) => {
  return (
    <div className="flex w-full min-w-0 flex-col gap-6 px-5 py-7 sm:px-8 lg:px-10">
      <Link
        href="/cohort"
        className="text-secondary-text hover:text-primary flex w-fit items-center gap-2 text-sm font-semibold transition-colors"
      >
        <ArrowLeft className="size-4" />
        Kelas Saya
      </Link>

      <Suspense fallback={<CohortDetailSkeleton />}>
        <CohortDetailShell params={params} />
      </Suspense>

      {children}
    </div>
  );
};

export default CohortDetailLayoutContainer;