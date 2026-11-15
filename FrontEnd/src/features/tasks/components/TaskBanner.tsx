import { ClipboardList } from "lucide-react";

const TaskBanner = () => {
  return (
    <div
      aria-hidden="true"
      className="from-primary to-secondary relative h-32 w-full overflow-hidden rounded-2xl bg-linear-to-r sm:h-40"
    >
      <div className="pointer-events-none absolute -top-16 -left-10 size-52 rounded-full border-28 border-white/10" />
      <div className="pointer-events-none absolute -right-8 -bottom-20 size-56 rounded-full bg-white/10" />
      <div className="pointer-events-none absolute top-6 left-2/3 size-16 rounded-full border-8 border-white/15" />
      <div className="pointer-events-none absolute bottom-5 left-1/3 size-6 rounded-full bg-white/20" />

      <div className="relative z-10 flex h-full flex-col items-center justify-center gap-2 px-4 text-center sm:gap-3">
        <div className="flex items-center justify-center gap-3 sm:gap-4">
          <ClipboardList className="h-auto w-12 shrink-0 text-white sm:w-16 lg:w-20" />
          <h1 className="text-3xl leading-tight font-bold text-white sm:text-5xl lg:text-6xl">
            Tugas Saya
          </h1>
        </div>
        <p className="max-w-xl text-sm text-white/80 sm:text-lg">
          Semua tugas dari kelas yang kamu ikuti, dalam satu daftar.
        </p>
      </div>
    </div>
  );
};

export default TaskBanner;
