import Image from 'next/image';
import { useContext } from 'react';
import formatNumber from '@/utils/formatNumber';
import Button from '../button/Button';
import GlobalAppContext from '../layout/context/GlobalAppContext';

type BannerProps = {
  githubStars: number | undefined;
};

function Banner({ githubStars }: BannerProps) {
  const { displayBanner, dismissBanner } = useContext(GlobalAppContext);

  return (
    <div
      className={`${
        displayBanner ? 'fixed' : 'hidden'
      } top-0 z-10 flex w-full animate-fade-in-down-short items-center justify-center gap-6 bg-gradient-to-br from-primary to-secondary py-3 opacity-0`}
    >
      <span className="text-sm font-medium text-white">
        Support Komiser by giving us a star on GitHub.
      </span>

      {githubStars && (
        <a
          href="https://github.com/tailwarden/komiser"
          target="_blank"
          rel="noreferrer"
          className="group"
        >
          <Button style="bulk-outline" size="md">
            <Image
              src="./assets/img/others/github-white.svg"
              width="18"
              height="16"
              alt="Github logo"
            />
            <span>Star Komiser</span>
            <div className="ml-2 -mr-6 flex h-full w-16 items-center justify-center gap-2 border-l border-white/10 bg-white/10">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="none"
                viewBox="0 0 16 16"
                className="group-hover:fill-warning-600 group-hover:text-warning-600"
              >
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="1.5"
                  d="M9.153 2.34l1.174 2.346c.16.327.586.64.946.7l2.127.354c1.36.226 1.68 1.213.7 2.186L12.447 9.58c-.28.28-.434.82-.347 1.206l.473 2.047c.374 1.62-.486 2.247-1.92 1.4l-1.993-1.18c-.36-.213-.953-.213-1.32 0l-1.993 1.18c-1.427.847-2.294.213-1.92-1.4l.473-2.047c.087-.386-.067-.926-.347-1.206L1.9 7.926c-.973-.973-.66-1.96.7-2.186l2.127-.354c.353-.06.78-.373.94-.7L6.84 2.34c.64-1.274 1.68-1.274 2.313 0z"
                ></path>
              </svg>
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
