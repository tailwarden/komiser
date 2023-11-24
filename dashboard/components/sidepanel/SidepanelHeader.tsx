import { ReactNode } from 'react';
import ArrowLeftIcon from '@components/icons/ArrowLeftIcon';
import HyperLinkIcon from '@components/icons/HyperLinkIcon';
import Button from '@components/button/Button';
import Avatar from '@components/avatar/Avatar';
import { Provider } from '@utils/providerHelper';

export type SidepanelHeaderProps = {
  title: string;
  subtitle?: string;
  href?: string;
  cloudProvider?: Provider;
  children?: ReactNode;
  closeModal: () => void;
  deleteAction?: () => void;
  goBack?: () => void;
  deleteLabel?: string;
};

function SidepanelHeader({
  title,
  subtitle,
  href,
  cloudProvider,
  children,
  closeModal,
  deleteAction,
  deleteLabel,
  goBack
}: SidepanelHeaderProps) {
  return (
    <div
      className={`flex flex-wrap-reverse items-center justify-between gap-6 sm:flex-nowrap ${
        subtitle && 'pt-2'
      }`}
    >
      {title && subtitle && (
        <div className="flex flex-wrap items-center gap-2 sm:flex-nowrap">
          {cloudProvider && <Avatar avatarName={cloudProvider} size={40} />}
          <div className="flex flex-col gap-0.5">
            <p className="font-['Noto Sans'] text-neutral-900 inline-flex w-48 items-center gap-2 truncate text-base font-medium leading-normal">
              {title}
              <a
                target="_blank"
                href={href}
                rel="noreferrer"
                className="hover:text-darkcyan-500"
              >
                <HyperLinkIcon />
              </a>
            </p>
            <p className="font-['Noto Sans'] text-neutral-500 text-xs font-normal">
              {subtitle}
            </p>
          </div>
        </div>
      )}

      {title && !subtitle && (
        <div className="flex flex-wrap items-center gap-4 sm:flex-nowrap">
          <button type="button" onClick={goBack}>
            <ArrowLeftIcon className="h-6 w-6" />
          </button>
          <div className="flex flex-col gap-0.5">
            <p className="font-['Noto Sans'] text-neutral-900 text-center text-xl font-semibold leading-loose">
              {title}
            </p>
          </div>
        </div>
      )}

      {children}

      <div className="flex flex-shrink-0 items-center gap-4">
        {deleteAction && (
          <Button style="delete" onClick={deleteAction}>
            {deleteLabel || 'Delete'}
          </Button>
        )}

        <Button style="secondary" onClick={closeModal}>
          Close
        </Button>
      </div>
    </div>
  );
}

export default SidepanelHeader;
