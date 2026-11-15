interface DashboardHeaderProps {
  total: number;
  mentorCount: number;
  memberCount: number;
}

const DashboardHeader = ({
  total,
  mentorCount,
  memberCount,
}: DashboardHeaderProps) => {
  return (
    <div className="flex flex-wrap items-end justify-between gap-4">
      <div className="space-y-1">
        <h2 className="text-2xl-bold text-primary-text">Kelas Saya</h2>
        <p className="text-sm-medium text-secondary-text">
          Semua kelas yang kamu kelola dan ikuti.
        </p>
      </div>

      <div className="text-secondary-text flex items-center gap-3 text-sm">
        <span>
          <strong className="text-primary-text">{total}</strong> Kelas
        </span>
        <span className="text-[#CBD5E1]">•</span>
        <span>{mentorCount} Mentor</span>
        <span className="text-[#CBD5E1]">•</span>
        <span>{memberCount} Member</span>
      </div>
    </div>
  );
};

export default DashboardHeader;
