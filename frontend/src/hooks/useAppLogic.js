import { useEffect, useState, useCallback } from "react";

export const useAppLogic = (router) => {
  const [showScrollTopButton, setShowScrollTopButton] = useState(false);
  const [lastClickY, setLastClickY] = useState(0);

  const smoothScroll = useCallback((targetPosition, duration) => {
    const startPosition = window.scrollY;
    const distance = targetPosition - startPosition;
    let startTime = null;

    const animation = (currentTime) => {
      if (!startTime) startTime = currentTime;
      const timeElapsed = currentTime - startTime;
      const progress = Math.min(timeElapsed / duration, 1);

      const ease = (t) => t * (2 - t);

      const scrollY = startPosition + distance * ease(progress);
      window.scrollTo(0, scrollY);

      if (progress < 1) {
        requestAnimationFrame(animation);
      }
    };

    requestAnimationFrame(animation);
  }, []);

  const handleScrollToClick = useCallback(
    (event) => {
      if (event.target.tagName === "A" || event.target.tagName === "BUTTON") return;

      const mouseY = event.clientY;
      const threshold = 200;

      if (Math.abs(mouseY - lastClickY) < threshold) return;

      setLastClickY(mouseY);

      const targetPosition = window.scrollY + mouseY - window.innerHeight / 2;
      smoothScroll(targetPosition, 800);
    },
    [lastClickY, smoothScroll]
  );

  const handleScrollToTop = useCallback(() => {
    smoothScroll(0, 800);
  }, [smoothScroll]);

  useEffect(() => {
    const handleRouteChange = () => {
      window.scrollTo(0, 0);
    };

    router.events.on("routeChangeComplete", handleRouteChange);

    document.addEventListener("click", handleScrollToClick);

    const handleScroll = () => {
      setShowScrollTopButton(window.scrollY > 300);
    };
    window.addEventListener("scroll", handleScroll);

    return () => {
      router.events.off("routeChangeComplete", handleRouteChange);
      document.removeEventListener("click", handleScrollToClick);
      window.removeEventListener("scroll", handleScroll);
    };
  }, [handleScrollToClick, router.events]);

  return {
    showScrollTopButton,
    handleScrollToTop,
    handleScrollToClick,
  };
};
