import { html, css, LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

type Level = 'info' | 'warn' | 'error';
type Code = number;
type Message = string;

type Notify = {
    message: Message;
    level: Level;
    code: Code;
}

type Notification = {
    id: number;
    message: Message;
    level: Level;
    code: Code;
}

@customElement('htmx-toasts')
export class Toasts extends LitElement {
    static styles = css`.toast { z-index: 9999; }`;

    @property({ type: Array })
    notifications = new Array<Notification>()

    connectedCallback(): void {
        super.connectedCallback();
        window.addEventListener('htmx-toasts:notify', ((e: CustomEvent<Notify>) => this._handleNotify(e)) as EventListener);
    }

    disconnectedCallback(): void {
        super.disconnectedCallback();
        window.removeEventListener('htmx-toasts:notify', ((e: CustomEvent<Notify>) => this._handleNotify(e)) as EventListener);
    }

    private _handleNotify(e: CustomEvent<Notify>): void {
        const notifcation = { id: e.timeStamp, ...e.detail }
        this.notifications.push(notifcation);
        setTimeout(() => this._remove(notifcation), 3000);
        this.requestUpdate();
    }

    private _remove(n: Notification): void {
        this.notifications = this.notifications.filter(i => i.id !== n.id)
    }

    protected render() {
        return html`
            <link href="https://cdn.jsdelivr.net/npm/daisyui/dist/full.css" rel="stylesheet" type="text/css">
            <div class="toast" role="status" aria-live="polite">
                ${this.notifications.map(n => html`
                    <div class="alert alert-${n.level}">
                        <span>${n.message}</span>
                        <button class="btn btn-sm btn-outline" @click=${() => this._remove(n)}>Close</button>
                    </div>
                `)}
            </div>
        `;
    }
}