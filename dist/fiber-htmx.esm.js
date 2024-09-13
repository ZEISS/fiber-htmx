var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __decorateClass = (decorators, target, key, kind) => {
  var result = kind > 1 ? void 0 : kind ? __getOwnPropDesc(target, key) : target;
  for (var i = decorators.length - 1, decorator; i >= 0; i--)
    if (decorator = decorators[i])
      result = (kind ? decorator(target, key, result) : decorator(result)) || result;
  if (kind && result) __defProp(target, key, result);
  return result;
};

// src/toasts.ts
import { html, css, LitElement } from "lit";
import { customElement, property } from "lit/decorators.js";
var Toasts = class extends LitElement {
  constructor() {
    super(...arguments);
    this.notifications = new Array();
  }
  connectedCallback() {
    super.connectedCallback();
    window.addEventListener("htmx-toasts:notify", (e) => this._handleNotify(e));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener("htmx-toasts:notify", (e) => this._handleNotify(e));
  }
  _handleNotify(e) {
    const notifcation = { id: e.timeStamp, ...e.detail };
    this.notifications.push(notifcation);
    setTimeout(() => this._remove(notifcation), 3e3);
    this.requestUpdate();
  }
  _remove(n) {
    this.notifications = this.notifications.filter((i) => i.id !== n.id);
  }
  render() {
    return html`
            <link href="https://cdn.jsdelivr.net/npm/daisyui/dist/full.css" rel="stylesheet" type="text/css">
            <div class="toast" role="status" aria-live="polite">
                ${this.notifications.map((n) => html`
                    <div class="alert alert-${n.level}">
                        <span>${n.message}</span>
                        <button class="btn btn-sm btn-outline" @click=${() => this._remove(n)}>Close</button>
                    </div>
                `)}
            </div>
        `;
  }
};
Toasts.styles = css`.toast { z-index: 9999; }`;
__decorateClass([
  property({ type: Array })
], Toasts.prototype, "notifications", 2);
Toasts = __decorateClass([
  customElement("htmx-toasts")
], Toasts);
export {
  Toasts
};
