"use client";

import { Card, CardContent } from "@/shared/components/ui/card";
import { BENEFITS } from "../data/data";
import { useBenefitsAnimation } from "@/features/home/hooks/useBenefits";
import { cn } from "cn";

export function Benefits() {
  const { sectionRef } = useBenefitsAnimation();

  return (
    <section
      ref={sectionRef}
      id="manfaat"
      className="from-background via-surface to-background scroll-mt-28 bg-linear-to-b px-4 py-[clamp(4rem,8vw,7rem)]"
    >
      <div className="container mx-auto">
        <div
          data-benefits-intro
          className="mx-auto mb-12 flex max-w-3xl flex-col items-center gap-4 text-center"
        >
          <h2 className="text-primary-text max-w-[60ch] text-[clamp(2rem,4vw,3.25rem)] leading-[1.15] font-extrabold tracking-tight text-balance">
            Lebih sedikit urusan teknis,
            <span className="text-secondary">
              {" "}
              lebih banyak waktu untuk belajar.
            </span>
          </h2>

          <p className="text-secondary-text max-w-[60ch] text-base leading-relaxed md:text-lg">
            Lentera membantu mentor dan anggota tetap terhubung, terorganisir,
            dan fokus pada perkembangan belajar.
          </p>
        </div>

        <div className="grid grid-cols-1 gap-5 md:grid-cols-3">
          {BENEFITS.map(({ icon: Icon, title, text }, index) => {
            const isFeatured = index === 1;

            return (
              <Card
                key={title}
                data-benefits-card
                className={cn(
                  "group relative isolate flex min-h-60 flex-col overflow-hidden rounded-3xl border p-7 shadow-none transition-[box-shadow,border-color] duration-300",
                  "motion-safe:hover:shadow-[0_18px_40px_-18px_rgba(10,23,43,0.25)]",
                  isFeatured
                    ? "border-primary bg-primary"
                    : index === 0
                      ? "border-secondary/10 bg-secondary-light"
                      : "border-border bg-background hover:border-secondary/30",
                )}
              >
                <div
                  aria-hidden="true"
                  className={cn(
                    "pointer-events-none absolute -top-16 -right-16 size-48 rounded-full",
                    isFeatured ? "bg-white/5" : "bg-secondary/5",
                  )}
                />

                <div className="relative flex items-start justify-between">
                  <div
                    className={cn(
                      "grid size-14 place-items-center rounded-2xl",
                      isFeatured
                        ? "bg-white/15 text-white"
                        : "text-secondary bg-white",
                    )}
                  >
                    <Icon
                      className="size-7 transition-transform duration-300 group-hover:scale-110"
                      strokeWidth={1.8}
                    />
                  </div>
                </div>

                <CardContent className="relative flex flex-1 flex-col gap-3 p-0">
                  <h3
                    className={cn(
                      "text-xl leading-snug font-bold tracking-tight",
                      isFeatured ? "text-white" : "text-primary-text",
                    )}
                  >
                    {title}
                  </h3>

                  <p
                    className={cn(
                      "text-sm leading-relaxed",
                      isFeatured ? "text-white/75" : "text-secondary-text",
                    )}
                  >
                    {text}
                  </p>
                </CardContent>

                <div
                  className={cn(
                    "absolute right-7 bottom-6 size-2 rounded-full",
                    isFeatured ? "bg-white/40" : "bg-secondary/25",
                  )}
                />
              </Card>
            );
          })}
        </div>
      </div>
    </section>
  );
}
