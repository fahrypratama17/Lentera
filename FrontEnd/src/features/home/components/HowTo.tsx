"use client";

import { Card, CardContent } from "@/shared/components/ui/card";
import { STEPS, VARIANTS } from "../data/data";
import { cn } from "cn";
import { useHowToAnimation } from "@/features/home/hooks/useHowTo";

export function HowTo() {
  const { sectionRef } = useHowToAnimation();

  return (
    <section
      ref={sectionRef}
      id="cara-menggunakan"
      className="bg-background scroll-mt-28 px-4 py-[clamp(4rem,8vw,7rem)]"
    >
      <div className="container mx-auto">
        <div
          data-howto-intro
          className="mx-auto mb-12 flex max-w-3xl flex-col items-center gap-4 text-center"
        >
          <h2 className="text-primary-text max-w-[25ch] text-[clamp(2rem,4vw,3.25rem)] leading-[1.15] font-extrabold tracking-tight text-balance">
            Satu alur belajar,
            <span className="text-secondary"> semua lebih terarah.</span>
          </h2>
          <p className="text-secondary-text max-w-[60ch] text-base leading-relaxed md:text-lg">
            Dari bergabung di kelas hingga menyelesaikan tugas, semuanya jadi
            lebih mudah. Cukup ikuti empat langkah sederhana berikut.
          </p>
        </div>

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {STEPS.map(({ n, icon: Icon, title, text, variant }) => {
            const style = VARIANTS[variant];

            return (
              <Card
                key={n}
                data-howto-card
                className={cn(
                  "group relative isolate flex h-full flex-col overflow-hidden rounded-3xl border p-6 pb-7 shadow-none transition-[box-shadow,border-color] duration-300 motion-safe:hover:shadow-[0_18px_40px_-18px_rgba(10,23,43,0.25)]",
                  style.card,
                )}
              >
                {variant === "active" && (
                  <div
                    aria-hidden="true"
                    className="pointer-events-none absolute -top-12 -right-12 size-40 rounded-full bg-white/5"
                  />
                )}

                <div className="relative flex items-start justify-between">
                  <div
                    className={cn(
                      "grid size-12 place-items-center rounded-2xl",
                      style.icon,
                    )}
                  >
                    <Icon className="size-6" strokeWidth={1.8} />
                  </div>
                </div>

                <CardContent className="relative flex flex-1 flex-col gap-2 p-0">
                  <h3
                    className={cn(
                      "text-lg leading-snug font-bold tracking-tight",
                      style.title,
                    )}
                  >
                    {title}
                  </h3>
                  <p className={cn("text-sm leading-relaxed", style.text)}>
                    {text}
                  </p>
                </CardContent>
              </Card>
            );
          })}
        </div>
      </div>
    </section>
  );
}
