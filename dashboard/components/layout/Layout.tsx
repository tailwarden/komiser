import { ReactNode } from 'react';
import Banner from '../banner/Banner';
import useBanner from '../banner/hooks/useBanner';
import Navbar from '../navbar/Navbar';

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  const { displayBanner, dismissBanner, githubStars } = useBanner();

  return (
    <>
      <Banner
        displayBanner={displayBanner}
        githubStars={githubStars}
        dismissBanner={dismissBanner}
      />
      <Navbar displayBanner={displayBanner} />
      <main
        className={`relative ${
          displayBanner ? 'mt-[145px]' : 'mt-[73px]'
        } min-h-screen bg-black-100 p-6 xl:px-8 2xl:px-24`}
      >
        {children}
      </main>
    </>
  );
}

export default Layout;
