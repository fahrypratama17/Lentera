const TaskSkeleton = () => {
  return (
    <div
      aria-busy="true"
      className="flex w-full min-w-0 flex-col gap-7 px-5 py-7 sm:px-8 lg:px-10"
    >
      <div className="h-32 w-full animate-pulse rounded-2xl bg-[#DCE6ED] sm:h-40" />

      <div className="grid grid-cols-2 gap-3 sm:grid-cols-3 lg:grid-cols-5">
        {Array.from({ length: 5 }).map((_, index) => (
          <div
            key={index}
            className="h-20 animate-pulse rounded-2xl border border-[#DCE6ED] bg-white"
          />
        ))}
      </div>

      <div className="h-14 w-full animate-pulse rounded-xl bg-[#DCE6ED]/60" />

      <div className="grid grid-cols-1 gap-4 xl:grid-cols-2">
        {Array.from({ length: 4 }).map((_, index) => (
          <div
            key={index}
            className="h-44 animate-pulse rounded-2xl border border-[#DCE6ED] bg-white"
          />
        ))}
      </div>
    </div>
  );
};

export default TaskSkeleton;
