export interface LoginForm {
    created_at: string
    expires_at: string
    id: string
    issued_at: string
    refresh: boolean
    request_url: string
    requested_aal: string
    type: string
    ui: Ui
    updated_at: string
}

export interface RegistrationForm {
    expires_at: string
    id: string
    issued_at: string
    request_url: string
    type: string
    ui: Ui
}

export interface Ui {
    action: string
    method: string
    nodes: Node[]
}

export interface Node {
    attributes: Attributes
    group: string
    messages: any[]
    meta: Meta
    type: string
}

export interface Attributes {
    disabled: boolean
    name: string
    node_type: string
    required?: boolean
    type: string
    value?: string
    autocomplete?: string
}

export interface Meta {
    label?: Label
}

export interface Label {
    context?: Context
    id: number
    text: string
    type: string
}

export interface Context { }