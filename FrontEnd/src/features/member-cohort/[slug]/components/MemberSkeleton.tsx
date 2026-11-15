const MemberSkeleton = () => {
  return (
    <div aria-busy="true" className="flex flex-col gap-6">
      {Array.from({ length: 2 }).map((_, group) => (
        <div key={group} className="flex flex-col gap-3">
          <div className="h-4 w-24 animate-pulse rounded bg-[#DCE6ED]" />
          <div className="overflow-hidden rounded-2xl border border-[#DCE6ED] bg-white">
            {Array.from({ length: 4 }).map((_, row) => (
              <div
                key={row}
                className="h-16 animate-pulse border-b border-[#DCE6ED] last:border-b-0"
              />
            ))}
          </div>
        </div>
      ))}
    </div>
  );
};

export default MemberSkeleton;
