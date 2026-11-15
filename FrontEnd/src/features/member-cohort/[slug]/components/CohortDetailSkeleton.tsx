const CohortDetailSkeleton = () => {
  return (
    <div aria-busy="true" className="flex flex-col gap-6">
      <div className="flex flex-col gap-4">
        <div className="h-52 animate-pulse rounded-2xl bg-[#DCE6ED]" />
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          {Array.from({ length: 3 }).map((_, index) => (
            <div
              key={index}
              className="h-20 animate-pulse rounded-2xl border border-[#DCE6ED] bg-white"
            />
          ))}
        </div>
      </div>
      <div className="h-11 animate-pulse rounded-xl bg-[#DCE6ED]/60" />
    </div>
  );
};

export default CohortDetailSkeleton;