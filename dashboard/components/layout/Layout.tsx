import { ReactNode } from 'react';
import Navbar from '../navbar/Navbar';

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  return (
    <>
      <Navbar />
      <main className="relative mt-[73px] min-h-screen bg-black-100 p-8 2xl:px-24">
        {children}
      </main>
    </>
  );
}

export default Layout;
