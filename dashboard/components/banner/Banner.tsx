import classNames from 'classnames';
import { ReactNode } from 'react';

export type BannerProps = {
  children: ReactNode;
  displayBanner: boolean;
  dismissBanner: () => void;
  style?: 'primary' | 'secondary';
};

function Banner({
  children,
  displayBanner,
  dismissBanner,
  style = 'primary'
}: BannerProps) {
  const bannerStyle = classNames(
    'top-0 z-10 flex w-full animate-fade-in-down-short items-center justify-center gap-6 bg-gradient-to-br py-3 opacity-0',
    {
      fixed: displayBanner,
      hidden: !displayBanner,
      'text-white from-darkcyan-500 to-darkcyan-700': style === 'primary',
      'text-black bg-white shadow-right': style === 'secondary'
    }
  );
  const buttonStyle = classNames(
    'absolute right-8 cursor-pointer rounded-lg p-3 transition-colors',
    {
      'text-white hover:bg-white/10 active:bg-gray-950': style === 'primary',
      'text-black hover:bg-black/10': style === 'secondary'
    }
  );

  return (
    <div className={bannerStyle}>
      {children}
      <button onClick={dismissBanner} className={buttonStyle}>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="none"
          viewBox="0 0 24 24"
        >
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
            d="M8 16l8-8M16 16L8 8"
          ></path>
        </svg>
      </button>
    </div>
  );
}

export default Banner;
