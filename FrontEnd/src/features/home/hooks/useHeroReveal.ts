"use client";

import { useLayoutEffect, useRef } from "react";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";

gsap.registerPlugin(ScrollTrigger);

export function useHeroReveal() {
  const heroRef = useRef<HTMLElement>(null);
  const leftRef = useRef<HTMLDivElement>(null);
  const rightRef = useRef<HTMLDivElement>(null);
  const contentRef = useRef<HTMLDivElement>(null);

  useLayoutEffect(() => {
    const hero = heroRef.current;
    const left = leftRef.current;
    const right = rightRef.current;
    const content = contentRef.current;

    if (!hero || !left || !right || !content) return;

    const mm = gsap.matchMedia();

    mm.add("(prefers-reduced-motion: no-preference)", () => {
      const ctx = gsap.context(() => {
        const tl = gsap.timeline({
          defaults: {
            ease: "none",
          },
          scrollTrigger: {
            trigger: hero,
            start: "top top",
            end: "+=100%",
            scrub: 1,
            pin: true,
            anticipatePin: 1,
          },
        });

        tl.to(
          left,
          {
            xPercent: -100,
            duration: 1,
          },
          0,
        )
          .to(
            right,
            {
              xPercent: 100,
              duration: 1,
            },
            0,
          )
          .fromTo(
            content,
            {
              opacity: 0.5,
              scale: 0.96,
            },
            {
              opacity: 1,
              scale: 1,
              duration: 0.7,
            },
            0.2,
          );
      }, hero);

      return () => ctx.revert();
    });

    mm.add("(prefers-reduced-motion: reduce)", () => {
      gsap.set([left, right], { display: "none" });
    });

    return () => mm.revert();
  }, []);

  return {
    heroRef,
    leftRef,
    rightRef,
    contentRef,
  };
}
