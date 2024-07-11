import io
import json
import oci
import json
import base64

from fdk import response

signer = oci.auth.signers.get_resource_principals_signer()

secret_client = oci.secrets.SecretsClient(config={}, signer=signer)

# Busca o Secret no Vault pelo Secret OCID

def read_secret_value(secret_client, secretid):
    secret_content = secret_client.get_secret_bundle(secretid).data.secret_bundle_content.content.encode('utf-8')
    decrypted_secret_content = base64.b64decode(secret_content).decode("utf-8")
    return decrypted_secret_content

def handler(ctx, data: io.BytesIO = None):

    try:
        cfg = ctx.Config()
        secretid = cfg["secretid"]
    except Exception as e:
        print('Missing function parameters', flush=True)
        raise

    password = read_secret_value(secret_client, secretid)
    data = json.loads(data.getvalue())

    if data["token"] == password:
        message = True
    else:
        message = False

    return response.Response(
        ctx, response_data=json.dumps(
            {"active": message}),
        headers={"Content-Type": "application/json"}
    )
