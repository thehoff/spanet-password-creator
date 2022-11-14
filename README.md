
SPAnet API password hashing
--------------------------

This is to create an encrypted version of the password input to be sent to the spanet API.

Included are binaries and the source code.


Binary
------
The binary takes a single input, your spaNet account password and encodes it so it can be used on the spanet API.

./spanet-password-generator {your account password}
{encrypted string}



Source Information
------------------

Source Android APK from Google Play Store and Decompile.

sources > com > spa2 > activities > app > utils > Utils.java

    public static String SHA256(String str, String str2, boolean z) {
        UnsupportedEncodingException e;
        String str3;
        if (str == null) {
            return "Error: secretKey is null";
        }
        try {
            SecretKeySpec secretKeySpec = new SecretKeySpec(str.getBytes(), "HmacSHA256");
            Mac mac = Mac.getInstance("HmacSHA256");
            mac.init(secretKeySpec);
            str3 = Base64.encodeToString(mac.doFinal(str2.getBytes()), 0).trim();
            if (!z) {
                return str3;
            }
            try {
                return URLEncoder.encode(str3, AsyncHttpResponseHandler.DEFAULT_CHARSET);
            } catch (UnsupportedEncodingException e2) {
                e = e2;
                e.printStackTrace();
                return str3;
            } catch (InvalidKeyException | NoSuchAlgorithmException unused) {
                return str3;
            }
        } catch (UnsupportedEncodingException e3) {
            e = e3;
            str3 = null;
        } catch (InvalidKeyException | NoSuchAlgorithmException unused2) {
            return null;
        }
    }

    public static String HeronEncrypt(String str) {
        return SHA256("1ld0gVand", str, false).replace("+", "1").replace("/", "2").replace("=", "3");
    }
    
