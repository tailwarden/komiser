import Image from 'next/image';
import { useContext } from 'react';
import formatNumber from '../../utils/formatNumber';
import Button from '../button/Button';
import LayoutContext from '../layout/context/LayoutContext';

type BannerProps = {
  githubStars: number | undefined;
};

function Banner({ githubStars }: BannerProps) {
  const { displayBanner, dismissBanner } = useContext(LayoutContext);

  return (
    <div
      className={`${
        displayBanner ? 'fixed' : 'hidden'
      } top-0 z-10 flex w-full items-center justify-center gap-6 bg-gradient-to-br from-primary to-secondary py-3`}
    >
      <span className="text-sm font-medium text-white">
        Support Komiser by giving us a star on GitHub.
      </span>

      {githubStars && (
        <a
          href="https://github.com/tailwarden/komiser"
          target="_blank"
          rel="noreferrer"
        >
          <Button style="bulk-outline" size="md">
            <Image
              src="./assets/img/others/github-white.svg"
              width="18"
              height="16"
              alt="Github logo"
            />
            <span>Star Komiser</span>
            <div className="ml-2 -mr-6 flex h-full w-12 items-center justify-center border-l border-white/10 bg-white/10">
              {formatNumber(githubStars)}
            </div>
          </Button>
        </a>
      )}

      <button
        onClick={dismissBanner}
        className="absolute right-8 cursor-pointer rounded-lg p-3 text-white transition-colors hover:bg-white/10 active:bg-black-900/10"
      >
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
