const DashboardHeader = () => {
  return (
    <div className="flex flex-wrap items-end justify-between gap-4">
      <div className="space-y-1">
        <h2 className="text-2xl-bold text-primary-text">Kelas Saya</h2>
        <p className="text-sm-medium text-secondary-text">
          Semua kelas yang kamu kelola dan ikuti.
        </p>
      </div>
    </div>
  );
};

export default DashboardHeader;
