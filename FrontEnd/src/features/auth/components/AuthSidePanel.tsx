import { ArrowRight, BookOpen, Lightbulb, UsersRound } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import Image from "next/image";
import AuthBackground from "@/features/auth/components/AuthBackground";
import { AuthSidePanelProps } from "@/features/auth/types/type";

const AuthSidePanel = ({ isRegister, onSwitch }: AuthSidePanelProps) => {
  return (
    <aside
      className={`absolute inset-y-0 left-1/2 z-10 hidden w-1/2 overflow-hidden bg-linear-to-br from-[#0D6684] via-[#176781] to-[#1F506D] text-white transition-transform duration-700 ease-in-out md:flex ${
        isRegister ? "-translate-x-full" : "translate-x-0"
      } `}
    >
      <AuthBackground />

      <div className="relative z-10 flex w-full flex-col items-center justify-center gap-7 px-10 text-center">
        <div className="flex size-28 items-center justify-center rounded-3xl border border-white/20 bg-white/95 p-3 shadow-xl shadow-black/10">
          <Image
            src="/logo.webp"
            alt="Logo aplikasi"
            width={96}
            height={96}
            className="h-full w-full object-contain"
            priority
          />
        </div>

        <div className="flex flex-col gap-3">
          <h2 className="text-3xl-bold text-white">
            {isRegister
              ? "Selamat Datang Kembali!"
              : "Belajar Lebih Baik Bersama"}
          </h2>

          <p className="text-md-medium leading-relaxed text-white/85">
            {isRegister
              ? "Sudah memiliki akun? Masuk untuk melanjutkan perjalanan belajarmu."
              : "Bergabung dengan cohort, jelajahi materi, dan berkembang bersama komunitasmu."}
          </p>
        </div>

        <div className="flex items-center gap-4">
          <div className="rounded-xl border border-white/20 bg-white/10 p-3 backdrop-blur-sm">
            <BookOpen size={22} />
          </div>
          <div className="rounded-xl border border-white/20 bg-white/10 p-3 backdrop-blur-sm">
            <UsersRound size={22} />
          </div>
          <div className="rounded-xl border border-white/20 bg-white/10 p-3 backdrop-blur-sm">
            <Lightbulb size={22} />
          </div>
        </div>

        <Button
          type="button"
          size="lg"
          onClick={onSwitch}
          className="text-sm-bold mt-3 h-12 min-w-44 cursor-pointer rounded-xl border border-white/20 bg-[#E6F3FB] px-6 text-[#1F506D] transition-all duration-200 hover:-translate-y-0.5 hover:bg-white active:translate-y-0"
        >
          {isRegister ? "Masuk" : "Daftar Sekarang"}
          <ArrowRight size={17} />
        </Button>
      </div>
    </aside>
  );
};

export default AuthSidePanel;
