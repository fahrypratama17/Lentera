"use client";

import { useLayoutEffect, useRef } from "react";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";

gsap.registerPlugin(ScrollTrigger);

export function useCtaAnimation() {
  const sectionRef = useRef<HTMLElement>(null);

  useLayoutEffect(() => {
    const section = sectionRef.current;
    if (!section) return;

    const mm = gsap.matchMedia();

    mm.add("(prefers-reduced-motion: no-preference)", () => {
      const ctx = gsap.context(() => {
        const elements = gsap.utils.toArray<HTMLElement>(
          "[data-cta-reveal]",
          section,
        );

        if (!elements.length) return;

        gsap.fromTo(
          elements,
          {
            autoAlpha: 0,
            y: 45,
          },
          {
            autoAlpha: 1,
            y: 0,
            duration: 0.8,
            stagger: 0.16,
            ease: "power3.out",
            clearProps: "opacity,visibility,transform",
            scrollTrigger: {
              trigger: section,
              start: "top 75%",
              once: true,
            },
          },
        );
      }, section);

      return () => ctx.revert();
    });

    return () => mm.revert();
  }, []);

  return { sectionRef };
}
