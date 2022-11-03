import { ReactNode } from "react";
import Navbar from "../navbar/Navbar";

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  return (
    <>
      <Navbar />
      <main className="relative min-h-screen bg-komiser-100 py-10 px-24">
        {children}
      </main>
    </>
  );
}

export default Layout;
