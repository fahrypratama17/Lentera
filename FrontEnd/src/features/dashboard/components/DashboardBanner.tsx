import { Plus } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import Image from "next/image";
import { DashboardBannerProps } from "@/features/dashboard/types/type";
import BannerBackground from "@/features/dashboard/components/BannerBackground";

const DashboardBanner = ({ onCreate }: DashboardBannerProps) => {
  return (
    <section className="relative overflow-hidden rounded-[24px] bg-[linear-gradient(110deg,#1F506D_0%,#0D6684_100%)] px-6 py-8 text-white sm:px-9 sm:py-10">
      <BannerBackground />

      <div className="relative z-10 flex items-center justify-between gap-6">
        <div className="flex max-w-xl flex-col items-start gap-4">
          <div className="space-y-2">
            <h1 className="text-2xl-bold sm:text-3xl-bold text-white">
              Selamat Datang Kembali!
            </h1>
            <p className="text-sm-medium max-w-lg leading-relaxed text-white/85">
              Kelola kelas, eksplorasi materi, dan tingkatkan progres belajarmu
              bersama komunitas.
            </p>
          </div>

          <Button
            type="button"
            onClick={onCreate}
            className="text-primary hover:bg-secondary-light mt-2 h-11 cursor-pointer rounded-xl bg-white px-5 font-bold"
          >
            <Plus className="size-4" />
            Buat Kelas Baru
          </Button>
        </div>

        <div className="hidden shrink-0 items-center justify-center rounded-[28px] border border-white/20 bg-white/10 p-5 backdrop-blur-sm md:flex">
          <Image
            src="/logo.webp"
            alt="Logo StudyTrack"
            width={125}
            height={125}
            className="size-28 object-contain"
          />
        </div>
      </div>
    </section>
  );
};

export default DashboardBanner;
