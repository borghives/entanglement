//entangle complexity to assure higher propability of intended operation
class Entanglement {
    constructor(html) {
        this.html = html
    }

    GetEntangleHeader() {
        let retheader = {}
        const sessionNonce  = GetSessionNonce()
        if (sessionNonce) {
            retheader['x-entanglement-nonce'] = sessionNonce
        }

        if (!this.html) {
            return retheader
        }
        
        const objToken = this.html.getAttribute("token")
        if (objToken) {
            retheader['x-entanglement-token'] = objToken
        }

        return retheader
    }

    SetCorrelation(key, value) {
        this.html.setAttribute("data-"+key, value)
    }

    GetCorrelation(key) {
        return this.html.getAttribute("data-"+key)
    }

    Update(properties) {
        if ('Token' in properties) {
            this.html.setAttribute('token', properties.Token)
            if ('Correlations' in properties) {
                for (const elemId in properties.Correlations) {
                    let entangle = Entanglement.FromHtml(document.getElementById(elemId))
                    if (entangle) {
                        let correlations = properties.Correlations[elemId]
                        for (const key in correlations) {
                            entangle.SetCorrelation(key, correlations[key])
                            const elements = document.querySelectorAll(`[entanglement="${key}"]`)
                            for (const element of elements) {
                                element.setAttribute("id", correlations[key])
                                element.setAttribute("entanglement", "")
                            }
                        }
                    }
                }
            }
        }

    }

    static FromHtml(startElement) {
        let element = startElement?.closest("entanglement-system")
        if (element) {
            return new Entanglement(element)
        }
    }
}


function GetSessionNonce() {
    const nonce = document.getElementsByTagName("entanglement-nonce")
    for (let i=0; i < nonce.length; i++) {
        return nonce[i].getAttribute("nonce")
    }
}