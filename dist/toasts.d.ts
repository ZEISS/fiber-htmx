import { LitElement } from 'lit';
type Level = 'info' | 'warn' | 'error';
type Code = number;
type Message = string;
type Notification = {
    id: number;
    message: Message;
    level: Level;
    code: Code;
};
export declare class Toasts extends LitElement {
    static styles: import("lit").CSSResult;
    notifications: Notification[];
    connectedCallback(): void;
    disconnectedCallback(): void;
    private _handleNotify;
    private _remove;
    protected render(): import("lit-html").TemplateResult<1>;
}
export {};
//# sourceMappingURL=toasts.d.ts.map