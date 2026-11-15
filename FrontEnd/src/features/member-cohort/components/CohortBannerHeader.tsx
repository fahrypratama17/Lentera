import { GraduationCap } from "lucide-react";

const CohortBannerHeader = () => {
  return (
    <div
      aria-hidden="true"
      className="from-primary to-secondary relative h-32 w-full overflow-hidden rounded-2xl bg-linear-to-r sm:h-40"
    >
      <div className="pointer-events-none absolute -top-16 -left-10 size-52 rounded-full border-28 border-white/10" />
      <div className="pointer-events-none absolute -right-8 -bottom-20 size-56 rounded-full bg-white/10" />
      <div className="pointer-events-none absolute top-6 left-2/3 size-16 rounded-full border-8 border-white/15" />
      <div className="pointer-events-none absolute bottom-5 left-1/3 size-6 rounded-full bg-white/20"
       />

       <div className="relative h-full z-10 flex flex-col items-center justify-center gap-3">
        <div className=" flex items-center justify-center gap-4">
            <GraduationCap className="text-white w-20 h-auto"/>
            <h1 className="text-white text-6xl leading-tight font-bold">
            Kelas Saya
            </h1>
        </div>
        <p className="text-white/80 text-xl">
          Semua kelas yang kamu ikuti dan kelola.
        </p>
      </div>
    </div>
  );
};

export default CohortBannerHeader;