"use client";

import { useLayoutEffect, useRef } from "react";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";

gsap.registerPlugin(ScrollTrigger);

export function useAboutAnimation() {
  const sectionRef = useRef<HTMLElement>(null);

  useLayoutEffect(() => {
    const section = sectionRef.current;
    if (!section) return;

    const mm = gsap.matchMedia();

    mm.add("(prefers-reduced-motion: no-preference)", () => {
      const ctx = gsap.context(() => {
        const intro = section.querySelector("[data-about-intro]");
        const description = section.querySelector("[data-about-description]");
        const visual = section.querySelector("[data-about-visual]");
        const stats = gsap.utils.toArray<HTMLElement>(
          "[data-about-stat]",
          section,
        );

        const tl = gsap.timeline({
          scrollTrigger: {
            trigger: section,
            start: "top 75%",
            once: true,
          },
        });

        if (intro) {
          tl.fromTo(
            intro,
            { autoAlpha: 0, y: 35 },
            {
              autoAlpha: 1,
              y: 0,
              duration: 0.7,
              ease: "power3.out",
            },
          );
        }

        if (description) {
          tl.fromTo(
            description,
            { autoAlpha: 0, y: 25 },
            {
              autoAlpha: 1,
              y: 0,
              duration: 0.65,
              ease: "power2.out",
            },
            "-=0.4",
          );
        }

        if (visual) {
          tl.fromTo(
            visual,
            { autoAlpha: 0, x: 55, scale: 0.95 },
            {
              autoAlpha: 1,
              x: 0,
              scale: 1,
              duration: 0.9,
              ease: "power3.out",
            },
            "-=0.6",
          );
        }

        if (stats.length) {
          tl.fromTo(
            stats,
            { autoAlpha: 0, y: 30 },
            {
              autoAlpha: 1,
              y: 0,
              duration: 0.6,
              stagger: 0.15,
              ease: "power3.out",
              clearProps: "opacity,visibility,transform",
            },
            "-=0.3",
          );
        }
      }, section);

      return () => ctx.revert();
    });

    return () => mm.revert();
  }, []);

  return { sectionRef };
}
