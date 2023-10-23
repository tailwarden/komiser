import { useState, useRef, useCallback, memo, SyntheticEvent } from 'react';
// eslint-disable-next-line import/no-extraneous-dependencies
import { toBlob } from 'html-to-image';
import Image from 'next/image';

import Modal from '@components/modal/Modal';
import Input from '@components/input/Input';
import settingsService from '@services/settingsService';
import Button from '@components/button/Button';
import { useToast } from '@components/toast/ToastProvider';
import Upload from '@components/upload/Upload';

// We define the placeholder here for convenience
// It's difficult to read when passed inline
const textAreaPlaceholder = `Example:
Steps to Reproduce
1. Describe the actions you took leading to the bug.
2. Include any specific settings or options you selected.

Expected Behavior
1. Explain what you expected to happen.
2. Detail how the feature or function should work.

Outcome
1. Describe what actually happened.
2. Include any error messages or unexpected behavior.`;

const useFeedbackWidget = (defaultState: boolean = false) => {
  const [showFeedbackModel, setShowFeedbackModal] = useState(defaultState);

  const FEEDBACK_MODAL_ID = 'feedback-modal';

  function openFeedbackModal() {
    setShowFeedbackModal(true);
  }

  function closeFeedbackModal() {
    setShowFeedbackModal(false);
  }

  function toggleFeedbackModal() {
    setShowFeedbackModal(!showFeedbackModel);
  }

  const screenshotModalFilter = (node: HTMLElement) =>
    !node.id?.startsWith(FEEDBACK_MODAL_ID);

  const FeedbackModal = () => {
    const [email, updateEmail] = useState('');
    const [description, updateDescription] = useState('');
    const [isTakingScreenCapture, setIsTakingScreenCapture] = useState(false);
    const [fileAttachement, setFileAttachement] = useState<File | null>(null);
    const [isSendingFeedback, setIsSendingFeedback] = useState(false);
    const { showToast } = useToast();

    async function takeScreenshot() {
      if (
        document.documentElement === null ||
        isSendingFeedback ||
        isTakingScreenCapture ||
        fileAttachement !== null
      ) {
        return;
      }
      setIsTakingScreenCapture(true);

      toBlob(document.documentElement, {
        cacheBust: true,
        filter: screenshotModalFilter
      })
        .then(async blob => {
          // setScreenshotBlob(blob);
          if (blob !== null) {
            const screenShotFile = new File(
              [blob],
              'Automated screen capture',
              {
                type: blob.type
              }
            );

            setFileAttachement(screenShotFile);
          }

          showToast({
            hasError: false,
            title: 'Screen capture',
            message:
              'A screenshot of your current page on Komiser has been captured and attached to your feedback.'
          });
        })
        .catch(err => {
          showToast({
            hasError: true,
            title: 'Screen capture failed',
            message:
              'The capture of your current page on Komiser couldn’t be saved. Please try again or upload a screenshot manually. Our support is also happy to help you!'
          });
        })
        .finally(() => {
          setIsTakingScreenCapture(false);
        });
    }

    function clearFeedbackForm() {
      setFileAttachement(null);
      updateDescription('');
      updateEmail('');
    }

    async function uploadFeedback(e: SyntheticEvent) {
      if (!isSendingFeedback) {
        try {
          setIsSendingFeedback(true);
          e.preventDefault();
          const formData = new FormData();

          formData.append('description', description);
          if (email) formData.append('email', email);
          if (fileAttachement && fileAttachement !== null)
            formData.append('image', fileAttachement);

          settingsService
            .sendFeedback(formData)
            .then(result => {
              showToast({
                hasError: false,
                title: 'Feedback sent',
                message:
                  result.Response ||
                  'Your insights are valuable to us. Thank you for helping us improve!'
              });
              clearFeedbackForm();
            })
            .catch(error => {
              showToast({
                hasError: true,
                title: 'Feedback',
                message: 'An Error happened. Maybe try again please!'
              });
            })
            .finally(() => {
              setIsSendingFeedback(false);
            });
        } catch {
          setIsSendingFeedback(false);
        }
      }
    }

    function uploadFile(attachement: File | null): void {
      setFileAttachement(attachement);
    }

    return (
      <>
        <Modal
          isOpen={showFeedbackModel}
          closeModal={() => closeFeedbackModal()}
          id={FEEDBACK_MODAL_ID}
        >
          <div className="w-[546px]">
            <h3 className="text-lg font-bold text-black-900">
              Describe your issue
            </h3>
            <p className="text-base text-black-400">
              By providing details of the issue you’ve encountered and outlining
              the steps to reproduce it, we’ll be able to give you better
              support.
            </p>
            <form onSubmit={uploadFeedback} className="mt-4">
              <Input
                disabled={isSendingFeedback}
                type="email"
                label="Your email"
                name="email"
                action={change => {
                  updateEmail(change.email);
                }}
                value={email}
                required
              />
              <textarea
                disabled={isSendingFeedback}
                rows={13}
                placeholder={textAreaPlaceholder}
                className="peer mt-4 w-full rounded bg-white px-4 pb-[0.75rem] pt-[1.75rem] text-sm text-black-900 caret-primary outline outline-[0.063rem] outline-black-200 focus:outline-[0.12rem] focus:outline-primary"
                onChange={event => updateDescription(event?.target?.value)}
                value={description}
                required
              />
              <div className="mt-4 h-[96px] max-h-[96px] w-full">
                <div className="flex h-full basis-1/2 items-stretch justify-stretch gap-2">
                  {fileAttachement === null && (
                    <>
                      <div
                        onClick={() => takeScreenshot()}
                        className="flex-1 grow cursor-pointer rounded border-2 border-black-170 py-5 text-center text-xs transition hover:border-[#B6EAEA] hover:bg-black-100"
                      >
                        <svg
                          className="mb-2 inline-block"
                          width="25"
                          height="24"
                          viewBox="0 0 25 24"
                          fill="none"
                          xmlns="http://www.w3.org/2000/svg"
                        >
                          <path
                            d="M2.25 9V6.5C2.25 4.01 4.26 2 6.75 2H9.25"
                            stroke="#0C1717"
                            strokeWidth="1.5"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                          />
                          <path
                            d="M15.25 2H17.75C20.24 2 22.25 4.01 22.25 6.5V9"
                            stroke="#0C1717"
                            strokeWidth="1.5"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                          />
                          <path
                            d="M22.25 16V17.5C22.25 19.99 20.24 22 17.75 22H16.25"
                            stroke="#0C1717"
                            strokeWidth="1.5"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                          />
                          <path
                            d="M9.25 22H6.75C4.26 22 2.25 19.99 2.25 17.5V15"
                            stroke="#0C1717"
                            strokeWidth="1.5"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                          />
                        </svg>

                        <p>
                          {isTakingScreenCapture
                            ? 'Taking screencapture...'
                            : 'Capture current screen'}
                        </p>
                      </div>
                      <span className="self-center justify-self-center text-sm text-black-400">
                        or
                      </span>
                    </>
                  )}
                  <div className="flex-1 grow">
                    <Upload
                      multiple={false}
                      fileOrFiles={fileAttachement}
                      handleChange={uploadFile}
                      onClose={() => setFileAttachement(null)}
                      disabled={
                        fileAttachement !== null ||
                        isSendingFeedback ||
                        isTakingScreenCapture
                      }
                      onTypeError={(err: string) =>
                        showToast({
                          hasError: true,
                          title: 'File upload failed',
                          message: err
                        })
                      }
                      onSizeError={(err: string) =>
                        showToast({
                          hasError: true,
                          title: 'File upload failed',
                          message: err
                        })
                      }
                    />
                  </div>
                </div>
              </div>
              <div className="mt-4 flex justify-between">
                <p className="text-xs text-black-400">
                  Need in depth assistance?
                  <br />
                  Email us at{' '}
                  <a
                    href="mailto:support@tailwarden.com"
                    className="text-primary"
                  >
                    support@tailwarden.com
                  </a>
                  .
                </p>
                <div className="flex gap-2">
                  <Button
                    size="xs"
                    disabled={isSendingFeedback}
                    style="ghost"
                    onClick={() => closeFeedbackModal()}
                  >
                    Cancel
                  </Button>
                  <Button type="submit" size="xs" disabled={isSendingFeedback}>
                    {isSendingFeedback ? `Sending...` : `Send Feedback`}
                  </Button>
                </div>
              </div>
            </form>
          </div>
        </Modal>
      </>
    );
  };

  return {
    openFeedbackModal,
    closeFeedbackModal,
    toggleFeedbackModal,
    FeedbackModal: memo(FeedbackModal)
  };
};

export default useFeedbackWidget;
