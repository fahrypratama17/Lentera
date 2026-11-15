import { ArrowRight, BookOpen, GraduationCap, UsersRound } from "lucide-react";
import { Button } from "@/shared/components/ui/button";

interface AuthSidePanelProps {
  isRegister: boolean;
  onSwitch: () => void;
}

const AuthSidePanel = ({ isRegister, onSwitch }: AuthSidePanelProps) => {
  return (
    <aside
      className={`from-primary to-secondary absolute inset-y-0 left-1/2 z-10 hidden w-1/2 overflow-hidden bg-linear-to-br text-white transition-transform duration-700 ease-in-out md:flex ${isRegister ? "-translate-x-full" : "translate-x-0"} `}
    >
      <svg
        viewBox="0 0 500 650"
        preserveAspectRatio="none"
        aria-hidden="true"
        className="pointer-events-none absolute inset-0 h-full w-full"
      >
        <circle cx="420" cy="90" r="170" fill="#E6F3FB" opacity="0.06" />

        <path
          d="M0 460 C130 360 310 560 500 410 L500 650 L0 650 Z"
          fill="#E6F3FB"
          opacity="0.08"
        />

        <path
          d="M0 540 C160 460 340 630 500 500 L500 650 L0 650 Z"
          fill="#E6F3FB"
          opacity="0.12"
        />
      </svg>

      <div className="relative z-10 flex w-full flex-col items-center justify-center gap-7 px-10 text-center">
        <div className="flex size-24 items-center justify-center rounded-3xl border border-white/20 bg-white/10 shadow-lg backdrop-blur-sm">
          <GraduationCap size={48} strokeWidth={1.5} />
        </div>

        <div className="flex flex-col gap-3">
          <h2 className="text-3xl-bold text-white">
            {isRegister
              ? "Selamat Datang Kembali!"
              : "Belajar Lebih Baik Bersama"}
          </h2>

          <p className="text-md-medium leading-relaxed text-white/80">
            {isRegister
              ? "Sudah memiliki akun? Masuk untuk melanjutkan perjalanan belajarmu."
              : "Bergabung dengan cohort, jelajahi materi, dan berkembang bersama komunitasmu."}
          </p>
        </div>

        <div className="flex items-center gap-4">
          <div className="rounded-xl bg-white/10 p-3">
            <BookOpen size={22} />
          </div>

          <div className="rounded-xl bg-white/10 p-3">
            <UsersRound size={22} />
          </div>

          <div className="rounded-xl bg-white/10 p-3">
            <GraduationCap size={22} />
          </div>
        </div>

        <Button
          type="button"
          size="lg"
          onClick={onSwitch}
          className="text-sm-bold text-primary hover:bg-secondary-light mt-3 h-12 min-w-44 cursor-pointer rounded-xl border border-white/30 bg-white px-6 transition-all duration-200 hover:-translate-y-0.5 active:translate-y-0"
        >
          {isRegister ? "Masuk" : "Daftar Sekarang"}
          <ArrowRight size={17} />
        </Button>
      </div>
    </aside>
  );
};

export default AuthSidePanel;
