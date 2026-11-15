import Image from "next/image";
import Link from "next/link";
import { ArrowLeft } from "lucide-react";

const notFound = () => {
  return (
    <main className="bg-surface relative flex min-h-dvh items-center justify-center overflow-hidden px-6 py-10">
      <div
        aria-hidden="true"
        className="bg-secondary-light/70 pointer-events-none absolute -top-24 -left-24 h-72 w-72 rounded-full blur-3xl"
      />
      <div
        aria-hidden="true"
        className="bg-secondary-light/60 pointer-events-none absolute -right-24 -bottom-24 h-80 w-80 rounded-full blur-3xl"
      />

      <section className="bg-background border-primary/10 relative z-10 flex w-full max-w-5xl flex-col items-center gap-10 rounded-[32px] border px-6 py-10 text-center shadow-[0_24px_80px_-30px_rgba(31,80,109,0.28)] md:flex-row md:px-10 md:py-12 md:text-left">
        <div className="flex w-full justify-center md:w-1/2">
          <div className="relative w-full max-w-85">
            <Image
              src="/mascot-confused.webp"
              alt="Mascot confused"
              width={420}
              height={420}
              priority
              className="h-auto w-full object-contain"
            />
          </div>
        </div>

        <div className="flex w-full flex-col items-center gap-5 md:w-1/2 md:items-start">
          <span className="text-sm-bold bg-secondary-light text-primary rounded-full px-4 py-2">
            Error 404
          </span>
          <div className="flex flex-col gap-3">
            <h1 className="text-4xl-bold text-primary-text">
              Halaman tidak ditemukan
            </h1>
            <p className="text-md-medium text-secondary-text max-w-xl leading-relaxed">
              Ups, sepertinya halaman yang kamu cari tidak tersedia, sudah
              dipindahkan, atau alamatnya salah. Yuk kembali ke halaman utama
              untuk melanjutkan perjalananmu.
            </p>
          </div>

          <Link
            href="/"
            className="text-sm-bold bg-primary hover:bg-secondary inline-flex h-12 items-center justify-center gap-2 rounded-xl px-6 text-white transition-all duration-200 hover:-translate-y-0.5 active:translate-y-0"
          >
            <ArrowLeft size={18} />
            Kembali ke Landing Page
          </Link>
        </div>
      </section>
    </main>
  );
};

export default notFound;
