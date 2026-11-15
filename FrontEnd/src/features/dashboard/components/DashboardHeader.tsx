import { BookOpen } from "lucide-react";

const DashboardHeader = () => {
  return (
    <div className="flex flex-wrap items-center justify-between gap-4">
      <div className="flex items-center gap-4">
        <div className="bg-secondary-light text-primary flex size-16 shrink-0 items-center justify-center rounded-2xl">
          <BookOpen className="size-9" strokeWidth={2} />
        </div>

        <div className="space-y-1">
          <h2 className="text-2xl-bold text-primary-text">Jelajahi Kelas</h2>
          <p className="text-sm-medium text-secondary-text">
            Temukan dan kelola kelas pembelajaranmu di sini.
          </p>
        </div>
      </div>
    </div>
  );
};

export default DashboardHeader;
