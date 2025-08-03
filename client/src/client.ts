import "./client.css";
import { JSX } from "./jsx";
import { audit } from "./audit";
import { menuInit } from "./menu";
import { modeInit } from "./mode";
import { flashInit } from "./flash";
import { linkInit } from "./link";
import { timeInit } from "./time";
import { autocompleteInit } from "./autocomplete";
import { modalInit } from "./modal";
import { tagsInit } from "./tags";
import { formInit } from "./form";
import { themeInit } from "./theme";
import { Message, socketInit } from "./socket";
import { appInit } from "./app";

declare global {
  // eslint-disable-line @typescript-eslint/consistent-type-definitions
  interface Window {
    rituals: {
      wireTime: (el: HTMLElement) => void;
      relativeTime: (el: HTMLElement) => string;
      autocomplete: (
        el: HTMLInputElement,
        url: string,
        field: string,
        title: (x: unknown) => string,
        val: (x: unknown) => string
      ) => void;
      setSiblingToNull: (el: HTMLElement) => void;
      initForm: (frm: HTMLFormElement) => void;
      flash: (key: string, level: "success" | "error", msg: string) => void;
      tags: (el: HTMLElement) => void;
      Socket: unknown;
    };
    audit: (s: string, ...args: any) => void; // eslint-disable-line @typescript-eslint/no-explicit-any
    JSX: (tag: string, attrs: unknown[]) => HTMLElement;
  }
}

export function init(): void {
  const [s, i] = formInit();
  const [wireTime, relativeTime] = timeInit();
  window.rituals = {
    wireTime: wireTime,
    relativeTime: relativeTime,
    autocomplete: autocompleteInit(),
    setSiblingToNull: s,
    initForm: i,
    flash: flashInit(),
    tags: tagsInit(),
    Socket: socketInit()
  };
  menuInit();
  modeInit();
  linkInit();
  modalInit();
  themeInit();
  window.audit = audit;
  window.JSX = JSX;
  appInit();
}

document.addEventListener("DOMContentLoaded", init);
