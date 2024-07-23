import {css, html, LitElement} from "lit-element";

class MultiSelect extends LitElement {

    static get properties() {
        return {

        };
    }

    constructor() {
        super();
    }

    render() {
        return html`
        <woop>WHOOP!</woop>
           `;
    }

    static get styles() {
        return css`

        `
    }
}

customElements.define('multi-select', MultiSelect)
