import requests
import gzip
import os
from collections import defaultdict
import time
from concurrent.futures import ThreadPoolExecutor


# Function to process a single file and extract potential coupons
def process_file(filepath):
    print(f"Processing file: {filepath}")
    potential_coupons = []
    try:
        with gzip.open(filepath, 'rt', encoding='utf-8') as f:
            for line in f:
                coupon = line.strip()
                if 8 <= len(coupon) <= 10:
                    potential_coupons.append(coupon)
    except Exception as e:
        print(f"Error processing {filepath}: {e}")
    return potential_coupons

# Main script execution
if __name__ == "__main__":
    start_time = time.time()

    urls = [
        "https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase1.gz",
        "https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase2.gz",
        "https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase3.gz"
    ]

    for url in urls:
        filename = url.split("/")[-1]
        print(f"Downloading {filename}...")
        response = requests.get(url, stream=True)
        if response.status_code == 200:
            with open(filename, 'wb') as f:
                for chunk in response.iter_content(chunk_size=8192):
                    f.write(chunk)
            print(f"Downloaded {filename}")
        else:
            print(f"Failed to download {filename}. Status code: {response.status_code}")


    all_potential_coupons = defaultdict(int)

    # List of downloaded files
    downloaded_files = [f for f in os.listdir() if f.endswith('.gz')]

    # Process files in parallel
    with ThreadPoolExecutor(max_workers=min(len(downloaded_files), os.cpu_count() or 1)) as executor:
        future_to_file = {executor.submit(process_file, filename): filename for filename in downloaded_files}

        for future in future_to_file:
            filename = future_to_file[future]
            try:
                potential_coupons = future.result()
                for coupon in potential_coupons:
                    all_potential_coupons[coupon] += 1
                print(f"Finished processing {filename}")
            except Exception as exc:
                print(f'{filename} generated an exception: {exc}')

    # Identify valid coupons (present in at least two files)
    valid_coupons = [coupon for coupon, count in all_potential_coupons.items() if count >= 2]

    # Write valid coupons to a new file
    output_filename = "valid_coupons.txt"
    with open(output_filename, 'w') as f:
        for coupon in valid_coupons:
            f.write(f"{coupon}\n")

    end_time = time.time()
    elapsed_time = end_time - start_time

    print("\n--- Processing Complete ---")
    print(f"Total potential coupons found: {len(all_potential_coupons)}")
    print(f"Total valid coupons found: {len(valid_coupons)}")
    print(f"Valid coupons saved to: {output_filename}")
    print(f"Total execution time: {elapsed_time:.2f} seconds")

