"use client";

import { useLayoutEffect, useRef } from "react";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";

gsap.registerPlugin(ScrollTrigger);

export function useHowToAnimation() {
  const sectionRef = useRef<HTMLElement>(null);

  useLayoutEffect(() => {
    const section = sectionRef.current;
    if (!section) return;

    const mm = gsap.matchMedia();

    mm.add("(prefers-reduced-motion: no-preference)", () => {
      const ctx = gsap.context(() => {
        const intro = section.querySelector<HTMLElement>("[data-howto-intro]");

        const cards = gsap.utils.toArray<HTMLElement>(
          "[data-howto-card]",
          section,
        );

        if (!intro || cards.length === 0) return;

        const tl = gsap.timeline({
          scrollTrigger: {
            trigger: section,
            start: "top 75%",
            once: true,
          },
        });

        tl.fromTo(
          intro,
          {
            autoAlpha: 0,
            y: 30,
          },
          {
            autoAlpha: 1,
            y: 0,
            duration: 0.65,
            ease: "power2.out",
          },
        );

        tl.fromTo(
          cards,
          {
            autoAlpha: 0,
            y: 50,
            scale: 0.95,
          },
          {
            autoAlpha: 1,
            y: 0,
            scale: 1,
            duration: 0.7,
            stagger: 0.18,
            ease: "power3.out",
            clearProps: "transform,opacity,visibility",
          },
          "-=0.2",
        );
      }, section);

      return () => ctx.revert();
    });

    return () => mm.revert();
  }, []);

  return { sectionRef };
}
