import { ReactNode } from "react";
import Navbar from "../navbar/Navbar";

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  return (
    <>
      <Navbar />
      <main className="relative min-h-screen bg-black-100 py-8 px-24">
        {children}
      </main>
    </>
  );
}

export default Layout;
