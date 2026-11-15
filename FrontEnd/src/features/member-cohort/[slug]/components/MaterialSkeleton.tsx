const MaterialSkeleton = () => {
  return (
    <div aria-busy="true" className="flex flex-col gap-3">
      {Array.from({ length: 4 }).map((_, index) => (
        <div
          key={index}
          className="h-24 animate-pulse rounded-2xl border border-[#DCE6ED] bg-white"
        />
      ))}
    </div>
  );
};

export default MaterialSkeleton;
