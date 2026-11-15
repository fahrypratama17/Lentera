"use client";

import { useState } from "react";
import { Card, CardContent } from "@/shared/components/ui/card";
import LoginForm from "./LoginForm";
import RegisterForm from "./RegisterForm";
import AuthSidePanel from "./AuthSidePanel";

const AuthCard = () => {
  const [isRegister, setIsRegister] = useState(false);
  const showLogin = () => setIsRegister(false);
  const showRegister = () => setIsRegister(true);

  return (
    <Card className="border-primary/10 bg-background relative w-full max-w-[1040px] gap-0 overflow-hidden rounded-[28px] border py-0 shadow-[0_24px_80px_-30px_rgba(31,80,109,0.3)] md:min-h-[680px]">
      <CardContent className="relative min-h-[620px] w-full gap-0 px-0 md:min-h-[680px]">
        <section
          aria-hidden={isRegister}
          inert={isRegister}
          className={`min-h-[620px] w-full items-center px-7 py-10 sm:px-12 md:absolute md:inset-y-0 md:left-0 md:flex md:min-h-0 md:w-1/2 md:px-12 md:py-12 md:transition-opacity md:duration-300 ${
            isRegister
              ? "hidden md:pointer-events-none md:opacity-0"
              : "flex md:opacity-100"
          } `}
        >
          <LoginForm onSwitch={showRegister} />
        </section>

        <section
          aria-hidden={!isRegister}
          inert={!isRegister}
          className={`min-h-[620px] w-full items-center px-7 py-10 sm:px-12 md:absolute md:inset-y-0 md:right-0 md:flex md:min-h-0 md:w-1/2 md:px-12 md:py-12 md:transition-opacity md:duration-300 ${
            isRegister
              ? "flex md:opacity-100"
              : "hidden md:pointer-events-none md:opacity-0"
          } `}
        >
          <RegisterForm onSwitch={showLogin} />
        </section>

        <AuthSidePanel
          isRegister={isRegister}
          onSwitch={isRegister ? showLogin : showRegister}
        />
      </CardContent>
    </Card>
  );
};

export default AuthCard;
