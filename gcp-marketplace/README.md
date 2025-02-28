# Steps to host Bloxone image on GCP Marketplace

Follow the steps mentioned [checklist](https://cloud.google.com/marketplace/docs/partners/vm)
1. Set up your Google Cloud environment so that you can distribute and display your product on Cloud Marketplace.
2. Review the pricing options, and select a pricing model.
3. Build your VM image.
4. Create your deployment package.
5. Add a label to track your product's associated consumption of Google Cloud resources.
6. Test your product end-to-end.
7. Submit your product to Cloud Marketplace. 
8. Maintain and monitor your product after it has launched.

## Build your VM Image
1. To upload new Bloxone [image](https://docs.infoblox.com/space/BloxOneInfrastructure/350027851/Google+Cloud+Portal+(GCP)+Deployment)
2. Create the [Licensed VM image](https://cloud.google.com/marketplace/docs/partners/vm/build-vm-image#create_a_licensed_vm_image)

## Create your deployment package
1. I followed [this](https://cloud.google.com/marketplace/docs/partners/vm/configure-terraform-deployment#simple-deployment) to generate terraform template for references.
2. Followed [this](https://cloud.google.com/marketplace/docs/partners/vm/configure-terraform-deployment#complex-deployment) to create final terraform module. 
3. Default, Custom (CLI deployment) terraform module available [here](./nios-x-cli-tf)
4. Custom (UI deployment) terraform module available [here](./nios-x-ui-tf)
5. Created zip from [dir](./nios-x-ui-tf) 
    ```
   zip nios-x-ui-tf.zip * -r
   ```
6. Upload to Bucket 
7. Update the Deployment package and click validate.

## Test your product
1. once validation is success. Click Preview for testing the deployment. 

##### TODO
1. HTTP_Proxy is not added.
2. Key/Value declaration not possible in UI, need evaluate.
3. Get the Deployment launch UI page reviewed.

### Limitation
1. map(string) type is not supported in UI based deployment package.